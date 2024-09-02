const ENTITY_REFERENCE_CLASS = 'entity-reference'

$(() => {
  $('header img').on('click', () => {
    open('https://mathxplains.com.br')
  })

  $('a').on('click', (e) => {
    const link = $(e.target)
    const ref = link.prop('href')
    
    if (!link.hasClass(ENTITY_REFERENCE_CLASS)) return
    
    e.preventDefault()
    const header = $('header')
    const headerHeight = header.outerHeight()
    const destId = ref.split('#')[1]
    const dest = $(`#${destId}`)

    $('html, body').animate({
      scrollTop: (dest.offset().top - headerHeight) - 10,
    }, 0)
  })
})