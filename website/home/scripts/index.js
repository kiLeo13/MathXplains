import { showModal } from './resources/modals.js'
import { fetchSelfUser, updateCache } from './resources/resources.js'
import { loadAppointments } from './resources/appointments.js'
import { USER_ROUTES } from './http/routes.js'

$(async () => {
  await updateUsername()
  await updateCache()
  loadAppointments()

  $('#new-appointment-button').on('click', () => {
    showModal()
  })
})

async function updateUsername() {
  const self = await fetchSelfUser()

  if (self.name) {
    setUsername(self.name)
  } else {
    location.href = USER_ROUTES.LOGIN
  }
}

function setUsername(name) {
  $('#username').text(name)
}