console.log("Script loaded!");

const rootTemplate = `
    <div class="container">
      <button class="btn-primary">Books</button>
      <button class="btn-primary">Movies</button>
      <button class="btn-primary">TV Shows</button>
      <button class="btn-primary">Games</button>
    </div>

    <div class="cards" id="cards">
    </div>
`;
const loginTemplate = `
    <div class="container">
      <div>
        <h2 class="login-title">Login</h2>
        <form id="login-form" class="login-form">
          <input type="text" id="username" name="username" placeholder="Username" required>
          <input type="password" name="password" id="password" placeholder="Password" required>
          <button type="submit" class="btn-secondary">Login</button>
        </form>
      </div>
    </div>
`;




async function submitLoginForm(form) {
  const formData = new FormData(form);
  console.log(formData);
  var object = {};
  formData.forEach((value, key) => object[key] = value);
  var json = JSON.stringify(object);
  try {
    const response = await fetch("/api/auth/login", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: json,
    });
    const data = await response.json();
    console.log(data);
    if (response.ok) {
      localStorage.setItem("token", data.token);
      router("/");
    } else {
      // show error
      console.log(data.error);
    }
  } catch (e) {
    console.log(e);
  }
}









const routes = [
  { path: '/', template: rootTemplate },
  { path: '/login', template: loginTemplate },
]

function router(path) {
  //window.history.pushState({}, '', path);
  routes.forEach(route => {
    if(route.path === path) {
      document.getElementById('app').innerHTML = route.template;

      if(path === '/login') {
        // need to add the event listener to form
        const loginForm = document.querySelector("#login-form");

        loginForm.addEventListener("submit", (event) => {
          event.preventDefault();
          console.log("caught the form submit!");
          submitLoginForm(loginForm);
        });
      }
    }
  });
}

async function appendCard(item, index) {
  const cards = document.getElementById("cards");
  let card = '<div class="card"><div class="card__img_container">';
  card += '<img src="'+item.media.image_url+'" alt="" class="card__img">';
  card += '</div>';
  card += '<a href="#followed" class="card__link"></a>';
  card += '<p class="card__txt">'+ item.media.title +'</p>';
  card += '</div>';
  cards.innerHTML += card;
}

async function populateTable() {
  console.log("populate");
  try {
    const response = await fetch("/api/user/media", {
      method: "GET",
      headers: {
        "Content-Type": "application/json",
        Authorization: "Bearer " + localStorage.getItem("token")
      },
    });
    const data = await response.json();
    console.log(data);
    if (response.ok) {
      data.forEach(appendCard);
    } else {
      // show error
      console.log(data.error);
    }
  } catch (e) {
    console.log(e);
  }
}

async function initialize() {
  if (localStorage.getItem('token')) {
    router('/');
    populateTable();
  } else {
    router('/login');
  }
}

initialize();
