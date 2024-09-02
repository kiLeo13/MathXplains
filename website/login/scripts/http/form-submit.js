import { setButtonLoading, isDisabled } from "../buttons.js"
import ROUTES from "./routes.js"

$(() => {
  const form = $('#login-form')

  form.on('submit', (e) => {
    e.preventDefault()

    if (isDisabled()) return

    const email = $('#email').val()
    const password = $('#password').val()

    setButtonLoading(true)

    handleRequest({
      email: email,
      password: password
    })
  })
})

async function handleRequest(credentials) {
  const resp = await fetch(ROUTES.CREATE_LOGIN, {
    method: 'POST',
    headers: {
      "Content-Type": "application/json"
    },
    body: JSON.stringify({
      email: credentials.email,
      password: credentials.password
    })
  })
  const json = await resp.json()
  
  if (resp.ok) {
    localStorage.setItem('accessToken', json.access_token)
    localStorage.setItem('idToken', json.id_token)
    localStorage.setItem('refreshToken', json.refresh_token)

    location.href = location.origin
  } else {
    setButtonLoading(false)
    displayError(json.message)
  }
}

function displayError(msg) {
  $('#error-message').css({
    opacity: 1
  }).text(msg)
}