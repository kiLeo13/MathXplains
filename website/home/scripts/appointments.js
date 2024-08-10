import { fetchAppointments } from './resources/resources.js';

$(() => {
  setInterval(loadAppointments, 30000)
})

/**
 * Updates all the appointments on screen, this function will always make
 * a new request to the API (if it fails, an empty array is used instead)
 */
export async function loadAppointments() {
  const data = await fetchAppointments()
  const appts = data.appointments

  displayCount(appts.length, data.max)
  
  if (appts.length === 0) {
    // toggleCreateButton(false)
    return
  }
  
  for (const appt of data.appointments) {
    appts.push(newAppointmentRow(appt))
  }
  
}

function displayCount(active, max) {
  $('#active-count').text(active)
  $('#total-amount').text(max)
}

function displayAppointments() {
  clearSection()

}

function displayEmpty() {
  clearSection()
  
  $('.appointments-wrapper').height(400)
  $('.empty-image').show()
}

function clearSection() {
  $('.appointment').remove()
}

function toggleCreateButton(flag) {
  const button = $('.new-appointment-button')

  if (flag) {
    button
      .removeClass('disabled')
      .prop("disabled", false)
  } else {
    button
      .addClass('disabled')
      .prop("disabled", true)
  }
}

function createAppointmentRow(appt) {

}