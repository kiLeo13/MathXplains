import ROUTES from "../http/routes.js"
import { loadAppointments } from "../resources/appointments.js"
import { closeModal, setSaveLoading, showError } from "../resources/modals.js"
import { getIdToken } from "../resources/resources.js"

export async function onSubmit(e) {
  console.log('hi?')
  e.preventDefault()

  const subj = $('#form-subject').val()
  const topic = $('#form-topic').val()
  const date = $('#form-date').val()
  const desc = $('#form-desc').val()

  if (!isFuture(date)) {
    showError('A data fornecida não pode ser hoje, ou estar no passado.')
    return
  }

  setSendingUI()
  const sent = await send({
    "subject": subj,
    "topic": topic,
    "description": desc,
    "scheduled_at": date,
  })

  if (sent.success) {
    closeModal()
    loadAppointments()
  } else {
    showError(sent.error)
  }
}

function setSendingUI() {
  setSaveLoading()
  $('.loader-icon').css("opacity", 1)
}

function isFuture(date) {
  const input = new Date(date)
  const today = new Date()
  today.setHours(0, 0, 0, 0)

  return input > today
}

async function send(form) {
  const resp = await fetch(ROUTES.CREATE_APPOINTMENT, {
    method: 'POST',
    headers: {
      "Content-Type": "application/json",
      "Authorization": `Bearer ${getIdToken()}`
    },
    body: JSON.stringify(form)
  })
  const json = await resp.json()

  return {
    success: resp.ok,
    appointment: json
  }
}