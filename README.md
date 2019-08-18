<h1 align="center"> flatten </h1> <br>
<p align="center">
    <img src="./assets/images/steam-roller.jpg">
</p>

<p align="center">
  Flatten multidimensional integer arrays. Built with Go.
</p>

<p align="center">
    <img src="./assets/images/go-logo.png" height="100" width="100">
</p>

## Table of Contents

- [Author](#author)
- [About](#about)
- [Usage](#usage)
- [Performace](#performance)

## Author
Barry T. Burch<br>

Barry is a digital native with over 20 years of experience in software (and hardware) design and engineering at:

<p align="middle">
    <img src="./assets/images/ti-logo-2.png" align="center" hspace="10">
    <img src="./assets/images/nec-logo-2.png" align="center" hspace="10">
    <img src="./assets/images/att-logo-2.jpeg" align="center" hspace="20">
    <img src="./assets/images/avaya-logo-2.png" width="100" align="center" hspace="10">
    <img src="./assets/images/sxm-logo.jpeg" width="100" align="center" hspace="10">
    <img src="./assets/images/gf-logo.jpeg" width="100" align="center" hspace="10">
</p>

barry@sbcglobal.net<br>
www.linkedin.com/in/barry-burch-digital-native<br>

## About

The flatten package was created to be submitted to Theorem LLC as the coding exercise portion of an initial screening for a position as a software engineer.


## Usage

Flatten takes a single input parameter: []interface{}. This Go data structure is suitable for holding a multidimensional
array where the elements can be of any type (including user defined types). E.g., given the following snippet of Go code:

var s = []interface{}{ []interface{}{1, 2, []interface{}{3, 4, 5, []interface{}{6, 7}} }}

s now represents the following multidimensional integer array: [[1 2 [3 4 5 [6 7]]]] and can be passed as a parameter
to Flatten.

The flatten package includes the helper function Compose that allows for a more intuitive way to build 
multidimensional arrays for input to Flatten, e.g.:

s := Compose(1, 2, Compose(3, 4, 5, Compose(6, 7)))

While Flatten accepts as input a multidimensional array whose elements can be any type, it will return an error
on the first element it finds that is not an integer, i.e. Flatten currently only supports integer arrays.

## Performance
The Flatten function in the flatten package uses a recursive algorithm that visits each element in the multidimensional array (the input) only once and therefore has a complexity of O(n), i.e. proportional to the number of array elements.

The Flatten function's recursion anchor is implicit in that it is not possible to supply an input that has an 
infinite number of sub-arrays (i.e. a multidimensional array that has an infinite number of dimensions). This means
that the recursion will always reach the anchor.

