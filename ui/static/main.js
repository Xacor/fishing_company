"use strict";

var myModalEl = document.getElementById('delete-boat-modal')
myModalEl.addEventListener('show.bs.modal', function (event) {
    let form = this.queryselector('form');
    form.action = event.relatedTarget.dataset.url;
})