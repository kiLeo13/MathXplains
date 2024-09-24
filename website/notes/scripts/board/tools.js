import { load } from "./board.js"
import ROUTES from "../http/routes.js"
import { loadingRefreshContent, loadingSave } from "../sidebar/buttons.js"

$(() => {
  const $board = $('#board-content')
  const $noteName = $('#note-name')

  $('#increase-font-size').on('click', () => {
    const size = parseInt($board.css('font-size'))
    $board.css({
      "font-size": `${size + 1}px`,
      "line-height": `${(size + 1) * 1.5}px`
    })
  })
  
  $('#decrease-font-size').on('click', () => {
    const size = parseInt($board.css('font-size'))
    $board.css({
      "font-size": `${size - 1}px`,
      "line-height": `${(size - 1) * 1.5}px`
    })
  })

  $('#save-file-button').on('click', async () => {
    const $board = $('#board-content')
    const noteId = $board.attr('note')

    if (!noteId || noteId.length === 0) {
      alert('Nada para salvar.')
      return
    }

    const text = $board.val()
    const fileName = $('#note-name').val()

    loadingSave()
    loadingRefreshContent()
    const resp = await saveNote(noteId, text, fileName)
    
    if (!resp.ok) {
      alert('Failed to save: ' + resp.body.message)
    }
    loadingSave(false)
    loadingRefreshContent(false)
  })
  
  $('#refresh-content-button').on('click', async () => {
    loadingSave()
    loadingRefreshContent()
    const id = $board.attr('note')
    const resp = await openNote(id)
    loadingSave(false)
    loadingRefreshContent(false)

    if (resp.status === 200) {
      load(id, resp.body.name, resp.body.content)
      return
    }
  
    alert('Error: ' + resp.body.message)
    if (resp.status === 404) {
      $el.remove()
    }
  })

  $noteName.on('keyup', () => {
    const id = $board.attr('note')
    const $noteItem = $('#note-' + id)

    $noteItem.find('.sidebar-file-title').text($noteName.val())
  })

  ///////////////////////////
  //                       //
  // KEYBAORD INTERACTIONS //
  //                       //
  ///////////////////////////
  $('#board-content, #note-name').on('keydown', (e) => {
    if (e.ctrlKey && (e.key === 'Enter' || e.key === 's')) {
      e.preventDefault()
      $('#save-file-button').trigger('click') 
    }

    if (e.ctrlKey && e.key === 'r') {
      e.preventDefault()
      $('#refresh-content-button').trigger('click')
    }
  })
})

async function openNote(id) {
  const resp = await ROUTES.OPEN_NOTE.send({
    path: {id: id},
    query: {profile: sessionStorage.getItem('profile')}
  })

  return {
    status: resp.status,
    body: await resp.json()
  }
}

async function saveNote(id, content, name) {
  const resp = await ROUTES.PUT_NOTE.send({
    headers: {Profile: sessionStorage.getItem('profile')},
    path: {id: id},
    body: {
      content: content,
      name: name
    }
  })

  return {
    ok: resp.ok,
    body: await resp.json()
  }
}