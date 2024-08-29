import { canResend, sendResendCooldown, showError } from "../resources/buttons.js"
import ROUTES from './routes.js'

$(() => {
  const resend = $('#resend-code')

  resend.on('click', async () => {
    if (!canResend()) return

    const email = $("#email").val()

    if (!email) {
      alert('Email é obrigatório!')
      return
    }

    waitCooldown()
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
  const json = await resp.json()

  if (!resp.ok) showError(json.error)
}