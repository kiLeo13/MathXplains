const MAX_CHARS = 500_000

$(() => {
  const $board = $('#board-content')
  const $charCount = $('.char-counter')
  const formatter = new Intl.NumberFormat('pt-BR')

  $board.on('input', () => {
    const length = $board.val().length

    if (length >= MAX_CHARS) {
      $charCount.css('background-color', 'rgba(255, 40, 40, 0.5)')
    } else {
      $charCount.css('background-color', 'rgba(0, 0, 0, 0.3)')
    }

    $charCount.text(`${formatter.format(length)} / ${formatter.format(MAX_CHARS)}`)
  })
})

export function load(id, name, content) {
  const $board = $('#board-content')
  const $name = $('#note-name')
  const $stateButtons = $('.note-state-button')
  $stateButtons.removeAttr('disabled')
  $board.val(content)
  $board.removeAttr('disabled')
  $board.attr('note', id)
  
  $name.val(name)
  $name.removeAttr('disabled')
  
  $('.no-board-img').hide()
}

export function unload() {
  const $board = $('#board-content')
  const $name = $('#note-name')
  const $stateButtons = $('.note-state-button')
  $stateButtons.attr('disabled', 'true')
  $board.val('')
  $board.attr('disabled', 'true')

  $name.val('')
  $name.attr('disabled')
  $('.no-board-img').show()
}