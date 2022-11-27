function showData() {
    let ShowName = document.getElementById('input-name').value;
    let ShowEmail = document.getElementById('input-email').value;
    let ShowPhone = document.getElementById('input-phone').value;
    let showSubject = document.getElementById('input-subject').value;
    let showMessage = document.getElementById('input-message').value;

    console.log(ShowName);
    console.log(ShowEmail);
    console.log(ShowPhone);
    console.log(showSubject);

    if (ShowName == '') {
        return alert('Nama harus diisi.!')
    } else if (ShowEmail == '') {
        return alert('Email Harus diisi!')
    } else if (ShowPhone == '') {
        return alert('Nomor telepon harus diisi')
    } else if (showSubject == ''){
        return alert('Subject harus diisi')
    }

    
    let emailReciever = 'ilahir66@gmail.com'

    let a = document.createElement('a');
    a.href = `mailto:${emailReciever}?subject:${showSubject}&body= Hello, My Name ${ShowName}, ${showMessage}` 

    a.click()
}