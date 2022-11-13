"use strict";

// var myModalEl = document.getElementById('delete-boat-modal')
// myModalEl.addEventListener('show.bs.modal', function (event) {
//     let form = this.queryselector('form');
//     form.action = event.relatedTarget.dataset.url;
// })
function DeleteBoatModal(event) {
    let form = this.querySelector('form');
    form.action = event.relatedTarget.dataset.url;

    let boatNameElement = document.getElementById('boat_name_modal');
    boatNameElement.innerHTML = event.relatedTarget.closest('tr').querySelector('#boat-name').textContent;
}


function DeleteBankModal(event) {
    let form = this.querySelector('form');
    form.action = event.relatedTarget.dataset.url;

    let bankNameElement = document.getElementById('bank_name_modal');
    bankNameElement.innerHTML = event.relatedTarget.closest('tr').querySelector('#bank-name').textContent;
}

function DeleteFishModal(event) {
    let form = this.querySelector('form');
    form.action = event.relatedTarget.dataset.url;

    let fishNameElement = document.getElementById('fish_name_modal');
    fishNameElement.innerHTML = event.relatedTarget.closest('tr').querySelector('#fish-name').textContent;
}

window.onload = function() {
    console.log("aaaaaaaa");
    var modalBoat = document.getElementById('delete-boat-modal')
    if (modalBoat != null) {
        modalBoat.addEventListener('show.bs.modal', DeleteBoatModal)
    }
   

    var modalBank = document.getElementById('delete-bank-modal')
    if (modalBank != null) {
        modalBank.addEventListener('show.bs.modal', DeleteBankModal)
    }

    var modalFish = document.getElementById('delete-fish-modal')
    if (modalFish != null) {
        modalFish.addEventListener('show.bs.modal', DeleteFishModal)
    }
    
    
}