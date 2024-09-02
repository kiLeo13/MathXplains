import { setButtonLoading, hideError, showError, canSubmit } from '../resources/buttons.js'
import ROUTES from './routes.js'

$(() => {
  const form = $('#confirmation-form')

  form.on('submit', (e) => {
    e.preventDefault()

    if (!canSubmit()) return

    const email = $('#email').val()
    const code = $('#code').val()

    setButtonLoading(true)
    handleRequest({
      email: email,
      code: code
    })
  })
})

async function handleRequest(credentials) {
  hideError()
  const resp = await fetch(ROUTES.CREATE_CONFIRMATION, {
    method: 'POST',
    headers: {
      "Content-Type": "application/json"
    },
    body: JSON.stringify(credentials)
  })
  setButtonLoading(false)
  
  if (resp.ok) {
    location.href = location.origin + '/login'
  } else {
    const json = await resp.json()
    showError(json.message)
  }
}