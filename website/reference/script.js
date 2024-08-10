$(() => {
  $('header img').on('click', () => {
    open(location.origin)
  })

  $('a').on('click', () => {
    const link = $('a')
    const id = link.attr('id')

    if (!id.startsWith('#'))
      return

    // https://stackoverflow.com/questions/6677035/scroll-to-an-element-with-jquery
  })
})