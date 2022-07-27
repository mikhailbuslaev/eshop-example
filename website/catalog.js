let routes = {};
let templates = {};

let catalogBody = document.getElementById('catalogBody');

async function fetchCards(url = '', quantity = 20) {
  const formData = new FormData();
  formData.append('quantity', quantity);
    const response = await fetch(url, {
      method: 'POST',
      body: formData
    });
    return response.json();
}
  

function renderCatalog() {
    let catalogBody = document.getElementById('catalogBody');
    fetchCards('http://localhost:1111/cards', 20)
    .then((data) => {
        for (const card of data.message) {
            let cardDiv = document.createElement('div');
            cardDiv.appendChild(
              document.createElement('h1')
            ).textContent = card.title;

            var img = document.createElement("img"); 
            img.src = card.picturepath;
            cardDiv.appendChild(img);

            cardDiv.appendChild(
              document.createElement('h3')
            ).textContent = card.description;

            catalogBody.appendChild(cardDiv);
        }
    });
};

renderCatalog();