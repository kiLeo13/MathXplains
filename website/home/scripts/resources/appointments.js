import { deleteAppointment } from '../modal/delete-appointment.js'
import { registerDeletion } from '../modal/swipe-deletion.js'
import { fetchAppointments, getSubjectById, getProfessorById } from './resources.js'

/**
 * Updates all the appointments on screen, this function will always make
 * a new request to the API (if it fails, an empty array is used instead)
 */
export async function loadAppointments() {
  const data = await fetchAppointments()
  const appts = data.appointments

  displayCount(data.active, data.max)
  
  if (data.active >= data.max) toggleCreateButton(false)
  
  setItems(appts)
}

function displayCount(active, max) {
  $('#active-count').text(active)
  $('#total-amount').text(max)
}

/**
 * Sets the appointments on screen.
 * 
 * This method does not append elements, instead, it overrides
 * the existing appointment list by first clearing the section and then
 * appending the elements.
 * 
 * If an empty array is provided, this operation is delegated to
 * {@link setDisplayStyle}.
 * 
 * @param {any} appts An array of appointment objects.
 */
async function setItems(appts) {
  clearSection()

  if (appts.length === 0) {
    setDisplayStyle()
    return
  }
  setDisplayStyle(false)

  let html = ''
  for (const appt of appts) {
    const el = await buildAppointment(appt)
    html += el
  }

  $('.items-container').append(html)
  $(`.delete-appt-button`).on('click', deleteAppointment)
  registerDeletion()
}

function setDisplayStyle(empty = true) {
  clearSection()

  const emptyImage = $('.empty-image')
  const itemsContainer = $('.items-container')
  const apptsWrapper = $('.appointments-wrapper')

  if (empty) {
    emptyImage.show()
    itemsContainer.css("align-items", "center")
    apptsWrapper.css("position", "absolute")
  } else {
    emptyImage.hide()
    apptsWrapper.css({"position": "absolute"})
    itemsContainer.css({
      "align-items": "stretch",
      "justify-content": "start"
    })
  }
}

function clearSection() {
  $('.appointment').remove()
}

function toggleCreateButton(flag) {
  const button = $('#new-appointment-button')

  if (flag) button.removeProp("disabled")
  else button.prop("disabled", true)
}

async function buildAppointment(appt) {
  return `
    <div class="appointment ${appt.is_active ? "" : "disabled"}" timestamp="${appt.created_at}">
      <div class="appt-status ${getActiveClassState(appt.is_active)}"></div>
      <div class="appt-heading">
        <div class="left-heading">
          <div class="appointment-heading-title">
            <span class="appt-title">${appt.topic}</span>
            <span class="appt-timestamp">${formatTimestamp(appt.created_at)}</span>
          </div>
          <div class="bottom">
            <span class="appt-specifications">${await resolveSpecs(appt.subject_id)}</span>
          </div> 
        </div>
        <div class="right-heading">
          <button id="appt-${appt.id}" class="delete-appt-button" ${isDeletable(appt.created_at) ? "" : "disabled"}>
            <svg class="icon_d90b3d" aria-hidden="true" role="img" xmlns="http://www.w3.org/2000/svg" width="24" height="24"
              fill="none" viewBox="0 0 24 24">
              <path
                d="M14.25 1c.41 0 .75.34.75.75V3h5.25c.41 0 .75.34.75.75v.5c0 .41-.34.75-.75.75H3.75A.75.75 0 0 1 3 4.25v-.5c0-.41.34-.75.75-.75H9V1.75c0-.41.34-.75.75-.75h4.5Z"
                class=""></path>
              <path fill-rule="evenodd"
                d="M5.06 7a1 1 0 0 0-1 1.06l.76 12.13a3 3 0 0 0 3 2.81h8.36a3 3 0 0 0 3-2.81l.75-12.13a1 1 0 0 0-1-1.06H5.07ZM11 12a1 1 0 1 0-2 0v6a1 1 0 1 0 2 0v-6Zm3-1a1 1 0 0 1 1 1v6a1 1 0 1 1-2 0v-6a1 1 0 0 1 1-1Z"
                clip-rule="evenodd" class=""></path>
            </svg>
          </button>
        </div>
      </div>
      <div class="appointment-description-container">
        <span class="appointment-description-content">${appt.description}</span>
      </div>
    </div>`
}

function isDeletable(time) {
  const now = Date.now()
  const date = new Date(time)
  const period = (now - date) / (1000 * 60 * 60)

  return period < 24
}

function getActiveClassState(active) {
  return active
    ? 'appt-status-active'
    : 'appt-status-not-active disabled'
}

function resolveSpecs(id) {
  const subj = getSubjectById(id)
  if (!subj) return "Desconhecido"

  const prof = getProfessorById(subj.professor_id)
  let spec = subj.name

  if (prof) spec += ` &dash; ${prof.name}`
  return spec
}

function formatTimestamp(iso) {
  const date = new Date(iso)
  const pad = (num) => num.toString().padStart(2, '0')

  const day = pad(date.getDate())
  const month = pad(date.getMonth() + 1)
  const year = date.getFullYear()
  const hours = pad(date.getHours())
  const minutes = pad(date.getMinutes())

  return `${day}/${month}/${year} - ${hours}:${minutes}`
}

function toEpochUTC(time) {
  return Math.floor(new Date(time).getTime() / 1000)
}