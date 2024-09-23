import { loadingCreate, loadingRefresh } from "../buttons.js"
import { createNote, requireProfile, loadNotes } from "../sidebar.js"

$(() => {
  const $createButton = $('#new-file-button')
  const $updateButton = $('#refresh-files-button')
  
  $createButton.on('click', async () => {
    const profile = sessionStorage.getItem('profile')
    const fileName = new Date().toJSON()
    
    if (profile == null || profile === '') {
      requireProfile()
      return
    }
  
    loadingCreate()
    loadingRefresh()
    await createNote({name: onlyNumbers(fileName)})
    await loadNotes()
    
    loadingCreate(false)
    loadingRefresh(false)
  })
  
  $updateButton.on('click', async () => {
    loadingCreate()
    loadingRefresh()
    
    await loadNotes()
    
    loadingCreate(false)
    loadingRefresh(false)
  })
})

function onlyNumbers(str) {
  return str.replace(/\D/g, '')
}