$(() => {
    $('#signup-redirect').on('click', () => {
      location.href = location.origin + '/signup'
    })
})