"use strict";

// var myModalEl = document.getElementById('delete-boat-modal')
// myModalEl.addEventListener('show.bs.modal', function (event) {
//     let form = this.queryselector('form');
//     form.action = event.relatedTarget.dataset.url;
// })


window.onload = function() {
    var myModalEl = document.getElementById('delete-bank-modal')
    myModalEl.addEventListener('show.bs.modal', function (event) {
        let form = this.querySelector('form');
        form.action = event.relatedTarget.dataset.url;

        let bankNameElement = document.getElementById('bank_name_modal');
        bankNameElement.innerHTML = event.relatedTarget.closest('tr').querySelector('#bank-name').textContent;
    })
    
}