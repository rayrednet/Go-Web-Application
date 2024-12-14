const apiUrl = 'http://localhost:8080/albums';
const addAlbumButton = document.getElementById('addAlbumButton');
const addAlbumForm = document.getElementById('addAlbumForm');
const overlay = document.getElementById('overlay');
const confirmationPopup = document.getElementById('confirmationPopup');
const closePopupButton = document.getElementById('closePopupButton');
const searchButton = document.getElementById('searchButton');
const searchInput = document.getElementById('searchAlbumId');
const albumsList = document.getElementById('albums');

// Function to search for an album by ID
const notFoundPopup = document.getElementById('notFoundPopup');
const closeNotFoundPopupButton = document.getElementById('closeNotFoundPopupButton');

// Function to search for an album by ID
async function searchAlbum() {
    const albumId = searchInput.value.trim();
    const messageElement = document.getElementById('searchMessage');

    // Clear previous message
    if (messageElement) {
        messageElement.remove();
    }

    if (!albumId) {
        alert('Please enter an album ID to search.');
        return;
    }

    try {
        const response = await fetch(`${apiUrl}/${albumId}`);
        if (response.ok) {
            const album = await response.json();

            // Create and display the "Showing album with ID" message
            const albumsSection = document.querySelector("main");
            const newMessageElement = document.createElement('p');
            newMessageElement.id = 'searchMessage';
            newMessageElement.textContent = `Showing album with ID: ${albumId}`;
            albumsSection.insertBefore(newMessageElement, albumsList);

            // Display the album
            albumsList.innerHTML = `
                <li>
                    ID: ${album.id} - "${album.title}" by ${album.artist} - $${album.price.toFixed(2)}
                </li>
            `;
        } else {
            if (!response.ok) {
                // Show "Album Not Found" popup
                notFoundPopup.style.display = 'block';
            }
        }
    } catch (error) {
        console.error('Error searching album:', error);
    }
}

// Function to close the "Album Not Found" popup
function closeNotFoundPopup() {
    notFoundPopup.style.display = 'none';
}

// Attach event listener to close the "Album Not Found" popup
closeNotFoundPopupButton.addEventListener('click', closeNotFoundPopup);


// Function to fetch and display all albums
async function fetchAlbums() {
    try {
        const response = await fetch(apiUrl);
        const albums = await response.json();
        albumsList.innerHTML = ''; // Clear the album list
        albums.forEach(album => {
            const li = document.createElement('li');
            li.textContent = `ID: ${album.id} - "${album.title}" by ${album.artist} - $${album.price.toFixed(2)}`;
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

// Attach event listeners
addAlbumButton.addEventListener('click', showForm);
overlay.addEventListener('click', closeForm);
document.getElementById('addAlbumForm').addEventListener('submit', addAlbum);
closePopupButton.addEventListener('click', closeConfirmationPopup);
searchButton.addEventListener('click', searchAlbum);

// Enable search by pressing Enter key
searchInput.addEventListener('keydown', (event) => {
    if (event.key === 'Enter') {
        event.preventDefault(); // Prevent form submission or default behavior
        searchAlbum(); // Trigger the search
    }
});

// Fetch and display albums when the page loads
fetchAlbums();
