$(() => {
    updatePasswordField()

    $('#login-redirect').on('click', () => {
      location.href = location.origin + '/login'
    })
})