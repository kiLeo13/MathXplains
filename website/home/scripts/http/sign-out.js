import CACHE from '../resources/cache.js'
import { ROUTES, USER_ROUTES } from './routes.js'

export function signOut() {
  console.log('called?')
  const accessToken = localStorage.getItem(CACHE.ACCESS_TOKEN)

  localStorage.removeItem(CACHE.ACCESS_TOKEN)
  localStorage.removeItem(CACHE.ID_TOKEN)
  localStorage.removeItem(CACHE.REFRESH_TOKEN)

  invalidate(accessToken)
}

async function invalidate(token) {
  const resp = await fetch(ROUTES.GLOBAL_SIGN_OUT, {
    method: "POST",
    headers: {
      "Content-Type": "application/json"
    },
    body: JSON.stringify({
      "access_token": token
    })
  })

  if (resp.ok) {
    location.href = USER_ROUTES.LOGIN
  } else {
    const json = await resp.json()
    alert('Failed: ' + json.message)
  }
}