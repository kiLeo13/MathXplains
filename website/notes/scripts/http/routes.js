const BASE_URL = 'https://api.mathxplains.com.br'

const ROUTES = {
  GET_NOTES: BASE_URL + '/notes?profile=',
  CREATE_NOTE: BASE_URL + '/notes',

  DELETE_NOTE: BASE_URL + '/notes/',
  PUT_NOTE: BASE_URL + "/notes/",
}

export default ROUTES