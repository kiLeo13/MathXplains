import { fetchSelfUser } from './resources/resources.js'
import { loadAppointments } from './appointments.js'

$(async () => {
  // updateUsername()
  loadAppointments()
})

async function updateUsername() {
  const self = fetchSelfUser()

  if (self.name) {
    setUsername(self.name)
  } else {
    location.href = `${location.origin}/login`
  }
}

function setUsername(name) {
  $('#user-name').val(name)
}