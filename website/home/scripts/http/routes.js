const BASE_URL = 'https://api.mathxplains.com.br'
export const USER_URL = 'https://mathxplains.com.br'

const ROUTES = {
  GET_SELF_USER: BASE_URL + '/users/@me',

  GET_APPOINTMENTS:   BASE_URL + "/appointments?active=true",
  CREATE_APPOINTMENT: BASE_URL + "/appointments",
  DELETE_APPOINTMENT: BASE_URL + "/appointments/",

  GET_SUBJECTS:   BASE_URL + "/subjects?available=true",
  GET_PROFESSORS: BASE_URL + "/professors?known=true",

  REFRESH_TOKEN: BASE_URL + "/users/refresh",
}

export default ROUTES