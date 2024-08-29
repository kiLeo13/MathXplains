import { USER_URL } from "./routes.js"
import ROUTES from "./routes.js"
import CACHE from "../resources/cache.js"

export function getIdToken() {
  return localStorage.getItem(CACHE.ID_TOKEN)
}

export function getAccessToken() {
  return localStorage.getItem(CACHE.ACCESS_TOKEN)
}

// TODO: Finish implementation
/**
 * This function attempts to refresh both `accessToken` and `idToken`.
 * 
 * If this operation fails
 * (the `refreshToken` is likely to have already expired too),
 * then all tokens in cache will be cleared and the user will be redirected
 * for logging in again.
 */
export async function refreshTokens() {
  const refreshToken = localStorage.getItem(CACHE.REFRESH_TOKEN)
  
  if (!refreshToken || refreshToken === "") {
    endSession()
    return
  }
  
  const resp = await post(ROUTES.REFRESH_TOKEN, {"refresh_token": refreshToken})

  if (!resp.ok) {
    endSession()
    return
  }

  const body = resp.body
  const accessToken = body.access_token
  const idToken = body.id_token


}

/**
 * This function immediately removes all `accessToken`, `idToken` and `refreshToken`
 * from cache and redirects the user to the login page.
 */
function endSession() {
  localStorage.removeItem(CACHE.ACCESS_TOKEN)
  localStorage.removeItem(CACHE.ID_TOKEN)
  localStorage.removeItem(CACHE.REFRESH_TOKEN)

  location.href = USER_URL  + '/login'
}

async function post(route, body) {
  const resp = await fetch(route, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(body)
  })
  const json = await resp.json()

  return {
    status: resp.status,
    ok: resp.ok,
    body: json
  }
}