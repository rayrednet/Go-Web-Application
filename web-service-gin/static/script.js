const apiUrl = 'http://localhost:8080/albums';
const addAlbumButton = document.getElementById('addAlbumButton');
const addAlbumForm = document.getElementById('addAlbumForm');
const overlay = document.getElementById('overlay');
const confirmationPopup = document.getElementById('confirmationPopup');
const closePopupButton = document.getElementById('closePopupButton');

// Function to fetch and display all albums
async function fetchAlbums() {
    try {
        const response = await fetch(apiUrl);
        const albums = await response.json();
        const albumsList = document.getElementById('albums');
        albumsList.innerHTML = '';
        albums.forEach(album => {
            const li = document.createElement('li');
            li.textContent = `${album.id}: "${album.title}" by ${album.artist} - $${album.price}`;
            albumsList.appendChild(li);
        });
    } catch (error) {
        console.error('Error fetching albums:', error);
    }
}

// Function to add a new album
async function addAlbum(event) {
    event.preventDefault();
    const title = document.getElementById('albumTitle').value;
    const artist = document.getElementById('albumArtist').value;
    const price = parseFloat(document.getElementById('albumPrice').value);

    const newAlbum = { title, artist, price };

    try {
        const response = await fetch(apiUrl, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(newAlbum),
        });

        if (response.ok) {
            fetchAlbums(); // Refresh the list after adding
            document.getElementById('addAlbumForm').reset();
            closeForm();

            // Show confirmation popup
            confirmationPopup.style.display = 'block';
        } else {
            console.error('Failed to add album');
        }
    } catch (error) {
        console.error('Error adding album:', error);
    }
}

// Show form popup
function showForm() {
    addAlbumForm.style.display = 'block';
    overlay.style.display = 'block';
}

// Close form popup
function closeForm() {
    addAlbumForm.style.display = 'none';
    overlay.style.display = 'none';
}

// Close confirmation popup
function closeConfirmationPopup() {
    confirmationPopup.style.display = 'none';
}

// Event listeners
addAlbumButton.addEventListener('click', showForm);
overlay.addEventListener('click', closeForm);
document.getElementById('addAlbumForm').addEventListener('submit', addAlbum);
closePopupButton.addEventListener('click', closeConfirmationPopup);

// Fetch and display albums when the page loads
fetchAlbums();
