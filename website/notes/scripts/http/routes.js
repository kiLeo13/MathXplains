const BASE_URL = 'https://api.mathxplains.com.br'

class NoteRequest {
  constructor(route, method) {
    this.route = BASE_URL + route
    this.method = method
  }

  /**
   * The following params are accepted:
   * 
   * - `data.body` is expected to be an object that will become a JSON
   * to be sent to the server as `Content-Type: application/json`.
   * 
   * - `data.path` is expected to be an object with name and value,
   * for example:
   * ```js
   * {
   *  "id": 4
   * }
   * ```
   * 
   * - `data.query` is expected to be an object, containing name and value,
   * for example:
   * ```js
   * {
   *  "profile": "root",
   *  "limit": 100
   * }
   * ```
   * 
   * - `data.headers` is expected to be an object, containing all the additional headers,
   * for example:
   * ```js
   * {
   *  "Authorization": "Bearer [token]",
   *  "Referer": "mywebsite.com"
   * }
   * ```
   * 
   * The examples above will assemble a request for the {@link ROUTES.OPEN_NOTE OPEN_NOTE} endpoint
   * like following:
   * `https://api.mathxplains.com.br/notes/4?profile=root&limit=100`
   * 
   * @param {object} data
   */
  async send(data) {
    let route = this.route

    // Resolving path params
    for (let name in data.path) {
      route = route.replace(`{${name}}`, data.path[name])
    }

    route += resolveQuery(data.query)

    return fetch(route, {
      method: this.method,
      body: JSON.stringify(data.body),
      headers: {
        "Content-Type": "application/json",
        ...data.headers
      }
    })
  }
}

const ROUTES = {
  LIST_NOTES:  new NoteRequest('/notes',      "GET"),
  OPEN_NOTE:   new NoteRequest('/notes/{id}', "GET"),
  CREATE_NOTE: new NoteRequest('/notes',      "POST"),

  DELETE_NOTE: new NoteRequest('/notes/{id}', "DELETE"),
  PUT_NOTE:    new NoteRequest('/notes/{id}', "PUT")
}

function resolveQuery(params = {}) {
  if (Object.keys(params).length === 0) return ''
  
  const query = new URLSearchParams(params)
  return '?' + query.toString()
}

export default ROUTES