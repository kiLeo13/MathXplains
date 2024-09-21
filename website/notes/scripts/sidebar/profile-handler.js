$(() => {
  const $profileForm = $('.sidebar-header')

  $profileForm.on('submit', saveProfile)
  $('#profile-in').on('keydown', (e) => {
    if (e.key === 'enter') $profileForm.trigger('submit')
  })
})

function saveProfile(e) {
  e.preventDefault()
  const $profile = $('#profile-in')

  sessionStorage.setItem('profile', $profile.val())
  $profile.va
  $profile.val('')

  // Remove focus
  $profile.trigger('blur')
}