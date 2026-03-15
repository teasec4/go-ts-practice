import { h, text, app } from "https://unpkg.com/hyperapp";
// import html from 'https://unpkg.com/hyperlit?module' html`...`
// https://unpkg.com/@hyperapp/html?module 

console.log("HyperApp script loaded successfully");

const ClickHandler = (state) => ({
  ...state,
  clicks: state.clicks + 1
});

const InputHandler = (state, event) => ({...state, name: event.target.value})

app({
  init: {name: 'Max Ko', clicks: 0},
  view: (state) => h('main', {}, [
    h('h1', {}, text('Hello ' + state.name)),
    h('p', {}, text('Clicks: ' + state.clicks)),
    h('button', {
      onclick: ClickHandler
    }, text('Click me')),
    h('div', {}, [
      h('p', {}, text('Name: ' + state.name)),
      h('input', {
        type: 'text',
        value: state.name,
        oninput: InputHandler
      }) 
    ]),
    
  ]),
  node: document.getElementById("app"),
})      