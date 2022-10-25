import Vue from "vue";

require('./bootstrap');
import LE from 'laravel-extend-assents'
require('./components-register');

LE.init().then(function (le) {
    //console.log(a)
    //console.log(LE)
    //console.log(this)
    let ap = le.Vue()
    //console.log(ap)
})


//console.log(bb)
//console.log(cc)
