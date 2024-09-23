import { loadingCreate, loadingRefresh } from "../buttons.js"
import { loadNotes } from "../sidebar.js"

$(() => {
  const $profileForm = $('.sidebar-header')

  $profileForm.on('submit', saveProfile)
  $('#profile-in').on('keydown', (e) => {
    if (e.key === 'enter') $profileForm.trigger('submit')
  })
})

async function saveProfile(e) {
  e.preventDefault()
  const $profile = $('#profile-in')

  sessionStorage.setItem('profile', $profile.val())
  $profile.val('')

  $profile.trigger('blur')
  
  loadingRefresh()
  loadingCreate()
  await loadNotes()
  loadingRefresh(false)
  loadingCreate(false)
}