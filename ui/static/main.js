"use strict";

// var myModalEl = document.getElementById('delete-boat-modal')
// myModalEl.addEventListener('show.bs.modal', function (event) {
//     let form = this.queryselector('form');
//     form.action = event.relatedTarget.dataset.url;
// })


window.onload = function() {
    console.log("aaaaaaaa");
    var myModalElboat = document.getElementById('delete-boat-modal')
    myModalElboat.addEventListener('show.bs.modal', function (event) {
        let form = this.querySelector('form');
        form.action = event.relatedTarget.dataset.url;

        let boatNameElement = document.getElementById('boat_name_modal');
        boatNameElement.innerHTML = event.relatedTarget.closest('tr').querySelector('#boat-name').textContent;
    })
    var myModalElBoat = document.getElementById('delete-boat-modal')
    myModalElBoat.addEventListener('show.bs.modal', function (event) {
        let form = this.querySelector('form');
        form.action = event.relatedTarget.dataset.url;

        let boatNameElement = document.getElementById('boat_name_modal');
        boatNameElement.innerHTML = event.relatedTarget.closest('tr').querySelector('#boat-name').textContent;
    })
}