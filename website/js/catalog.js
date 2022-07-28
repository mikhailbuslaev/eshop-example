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
            cardDiv.className = "cardDiv";

            let title = document.createElement("h4"); 
            title.className = "cardTitle";
            title.textContent = card.title;
            cardDiv.appendChild(title);

            let img = document.createElement("img"); 
            img.src = card.picturepath;
            img.className = "cardImg";
            cardDiv.appendChild(img);

            let description = document.createElement('p');
            description.className = "cardPrice";
            description.textContent = card.price;
            cardDiv.appendChild(description);

            catalogBody.appendChild(cardDiv);
        }
    });
};

renderCatalog();