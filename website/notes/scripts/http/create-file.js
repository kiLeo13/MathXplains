import { createNote, requireProfile } from '../sidebar/resources.js'

$(() => {
  const $createButton = $('#new-file-button')

  $createButton.on('click', async () => {
    const profile = sessionStorage.getItem('profile')
    const fileName = new Date().toJSON()
    
    if (profile == null || profile === '') {
      requireProfile()
      return
    }

    $createButton.attr('disabled', 'true')
    const note = await createNote({name: onlyNumbers(fileName)})

    setTimeout(() => {
      $createButton.removeAttr('disabled')
    }, 1000)
  })
})

function onlyNumbers(str) {
  return str.replace(/\D/g, '')
}