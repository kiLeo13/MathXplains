import { ROUTES } from './routes.js'
import { getIdToken } from '../resources/resources.js'
import { loadAppointments } from '../resources/appointments.js'

export async function deleteAppointment(e) {
  console.log(JSON.stringify(e));
  const $el = $(e.currentTarget)
  const rawId = $el.attr('id')

  if (!rawId) {
    alert(`Failed? Could not resolve id of element: \n${$el.html()}`)
    return
  }

  const id = rawId.split('-')[1]
  const resp = await sendDelete(id)

  if (resp.ok) {
    loadAppointments()
  } else {
    alert('Não foi possível deletar o agendamento: ' + resp.message)
  }
}

async function sendDelete(id) {
  const resp = await fetch(ROUTES.DELETE_APPOINTMENT + id, {
    method: "DELETE",
    headers: {
      "Authorization": "Bearer " + getIdToken()
    }
  })
  
  if (resp.ok) {
    return {ok: true}
  } else {
    const json = await resp.json()
    return {
      ok: false,
      message: json.message
    }
  }
}