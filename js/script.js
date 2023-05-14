// Copyright (c) 2023 Lamees Hemdan All rights reserved
//
// Created by: Lamees Hemdan
// Created on: May 2023
// This file contains the JS functions for index.html

"use strict"

function myButtonClicked () {
  // this program creates a random number, either positive or negative.
  let randomNumber = 0.00
  // input
  const positiveButtonChecked = document.getElementById("positive").checked

  // process
  if (positiveButtonChecked == true) {
    // positive
    randomNumber = Math.floor(Math.random() * 6)
  } else {
    // negative
    randomNumber = Math.floor(Math.random() * -6)
  }

  // output
  document.getElementById('random-number').innerHTML = randomNumber
}