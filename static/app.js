import { h, text, app } from "https://unpkg.com/hyperapp";
// import { main, input, button, ul, li } from "https://unpkg.com/@hyperapp/html?module";
import { api } from "./api.js";

console.log("HyperApp script loaded successfully");

const fetchUsersFx = (dispatch) => {
  api
    .getUsers()
    .then((data) => dispatch(UsersLoaded, data))
    .catch((err) => dispatch(ApiError, err));
};

const LoadUsers = (state) => [state, fetchUsersFx];

const ApiError = (state, error) => {
  console.error(error);
  return state;
};

const UsersLoaded = (state, users) => ({
  ...state,
  users,
});

const InputNewItem = (state, event) => ({
  ...state,
  newname: event.target.value,
});

const createUserFx = (dispatch, props) => {
  api
    .createUser(props.name)
    .then((user) => dispatch(UserCreated))
    .catch((err) => dispatch(ApiError, err));
};

// [effect, props]
const CreateUser = (state) => [state, [createUserFx, { name: state.newname }]];

const UserCreated = (state) => [{ ...state, newname: "" }, fetchUsersFx];

app({
  init: [{ newname: "", users: [] }, fetchUsersFx],
  view: (state) =>
    h("main", {}, [
      h("input", {
        value: state.newname || "",
        oninput: InputNewItem,
      }),
      h("button", { onclick: CreateUser }, text("Add")),
      h(
        "ul",
        {},
        (state.users || []).map((user) => h("li", {}, text(user.name))),
      ),
    ]),
  node: document.getElementById("app"),
});
