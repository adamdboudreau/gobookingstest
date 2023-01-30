
    // https://sweetalert2.github.io/#examples
    // icon - error, info, success, warning
    function notifyModal(title, text, icon, confirmationButtonText, footer) {
        Swal.fire({
            icon: icon,
            title: title,
            html: text,
            confirmButtonText: confirmationButtonText,
            footer: footer
        })
    }
    
    function Prompt() {
        let toast = function (c) {
            const {
                msg = "",
                icon = "success",
                position = "top-end",
            } = c;
            const Toast = Swal.mixin({
            toast: true,
            title: msg,
            position: position,
            icon: icon,
            showConfirmButton: false,
            timer: 3000,
            timerProgressBar: true,
            didOpen: (toast) => {
                toast.addEventListener('mouseenter', Swal.stopTimer)
                toast.addEventListener('mouseleave', Swal.resumeTimer)
            }
            })

            Toast.fire({})
        }

        let success = function (c) {
            const {
                title = "",
                text = "",
                confirmationButtonText = "Ok",
                footer = "",
            } = c;

            Swal.fire({
                icon: "success",
                title: title,
                html: text,
                confirmButtonText: confirmationButtonText,
                footer: footer,
            })
        }

        let error = function (c) {
            const {
                title = "",
                text = "",
                confirmationButtonText = "Ok",
                footer = "",
            } = c;

            Swal.fire({
                icon: "error",
                title: title,
                html: text,
                confirmButtonText: confirmationButtonText,
                footer: footer,
            })
        }

        async function custom(c) {
            const {
                msg = "",
                title = "",
                text = "",
                confirmButtonText = "Submit",
            } = c;

            const {value: formValues } = await Swal.fire({
                icon: "success",
                title: title,
                html: msg,
                backdrop: false,
                focusConfirm: false,
                showCancelButton: true,
                confirmButtonText: confirmButtonText,
                willOpen: () => {
                    const elem = document.getElementById("reservation-dates-modal");
                    const rp = new DateRangePicker(elem, {format: 'yyyy-mm-dd'});
                },
                preConfirm: () => {
                    return [
                        document.getElementById('start').value,
                        document.getElementById('end').value,
                    ]
                },
                didOpen: () => {
                    return [
                        document.getElementById("start").removeAttribute('disabled'),
                        document.getElementById("end").removeAttribute('disabled')
                    ]
                }
            })
            if (formValues) {
                Swal.fire(JSON.stringify(formValues));
            }
        }

        return { 
            toast: toast, 
            success: success, 
            error: error,
            custom: custom, 
        }
    }