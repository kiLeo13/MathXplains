import { canResend, sendResendCooldown, showError } from "../resources/buttons.js"
import ROUTES from './routes.js'

$(() => {
  const resend = $('#resend-code')

  resend.on('click', async () => {
    if (!canResend()) return

    const email = $("#email").val()

    if (!email) {
      showError('Email é obrigatório!')
      return
    }

    sendResendCooldown()
    handleResend(email)
  })
})

async function handleResend(email) {
  const resp = await fetch(ROUTES.RESEND_CONFIRMATION, {
    method: 'POST', 
    body: JSON.stringify({ email: email }),
    headers: {
      "Content-Type": "application/json"
    }
  })
  
  if (!resp.ok) {
    const json = await resp.json()
    showError(json.message)
  }
}