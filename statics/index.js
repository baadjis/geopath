//////////////////////////////////////////// const and var/////////////////////////////////

var close_span = document.getElementsByClassName("close")[0];

let new_path_modal = document.getElementById("new-path-modal")
let coordinates_container = document.getElementById("coordinates-container")
let path_data_container = document.getElementById("path-data-container")
let Distances;
let Durations;
let path_names;
let new_path_name = document.getElementById("new-path-name")

let markers = []

var geojsonMarkerOptions = {
    radius: 8,
    fillColor: "#ff7800",
    color: "#000",
    weight: 1,
    opacity: 1,
    fillOpacity: 0.8
};

// When the user clicks on <span> (x), close the add path form modal
close_span.onclick = function () {
    new_path_modal.style.display = "none";
}


/// add option to select input to select a paths
async function addOptions() {

    const path_names_response = await fetch("/getPathNames")
    path_names = await path_names_response.json()



    path_names.forEach(element => {
        var opt = document.createElement('option');
        opt.text = element;
        opt.value = element;
        path_select.add(opt)

    });


}

// select input for selecting path
let path_select = document.getElementById("path-select")

// initialize the map
var map = L.map('map').setView([42.35, -71.08], 3);
let mypath;
/////////////////////map legend/////////////////////////////
let legend_div;

// load a tile layer
L.tileLayer('https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png', {
    attribution: '&copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors'
}).addTo(map);

var legend = L.control({ position: "bottomleft" });
legend.onAdd = function (map) {
    legend_div = L.DomUtil.create("div", "legend");
    return legend_div
}
legend.addTo(map)



//////////////////////////add a div to describe paths///////////////////////
//   <h1>pathname</h1> ////////////
/////////// distance/////////////
//////////////////duration///////////
 
/////////////////////////////////////////////////////////////////////:////////
function addPathDataDivs() {
    selected_path = path_select.value
    if (selected_path === "all") {
        path_names.map((path) => {
            path_data_container.innerHTML += `<div class="path-data"><h3>${path}</h3>
                         <span>Distance: ${Distances[`${path}`]}</span><br>
                         <span>Duration: ${Durations[`${path}`]} nanosecondes</span><br>
                         </div><br>`
        })
    } else {
        path_data_container.innerHTML += `<div class="path-data"><h3>${selected_path}</h3>
                         <p><span>Distance: ${Distances[`${selected_path}`]}</span><br>
                         <span>Duration: ${Durations[`${selected_path}`]} nanosecondes </span><br>
                         </p></div><br>`
    }
}

//////////////// update map  whenever  we change the select value//////////////////
function updatePath() {
    map.removeLayer(mypath)
    for (marker of markers) {
        map.removeLayer(marker)
    }

    markers = []
    legend_div.innerHTML = ''
    path_data_container.innerHTML = ''
    addGeoJson()




}

path_select.onchange = updatePath


//////////////////////////////////////display the form to create new path/////////////////////
function displayNewPathForm() {

    new_path_modal.style.display = "block"
}

let add_new_path_form_button =document.getElementById("add-new-path-form-button")
 
add_new_path_form_button.onclick=displayNewPathForm


////////////////////////////////////////////display new input to add other coordinates//////
function displayNewpositionForm() {

    var mydiv = document.createElement('div')
    mydiv.class = "coordinates"
    mydiv.innerHTML = '<input type="text"  placeholder="longitude" onkeypress="return isNumberKey(event)" required/> ' +
        '<input type="text" placeholder="latitude"   onkeypress="return isNumberKey(event)" required/> ' +
        '<input type="text" placeholder="altitude"  onkeypress="return isNumberKey(event)" required/> ' +
        '<input type="text" placeholder="timestamp" onkeypress="return isNumberKey(event)" required/> ';
    mydiv.outerHTML = '<br>';
    coordinates_container.appendChild(mydiv)


}
let display_new_position_button =document.getElementById("add-new-position-button")
display_new_position_button.onclick= function(event){
    event.preventDefault();
    displayNewpositionForm()

}

////////////////////////////ensure to type number or dot on coordinates input ///////////////////////
function isNumberKey(event) {
    return (event.charCode >= 48 && event.charCode <= 57) ||
        event.charCode == 46 || event.charCode == 0 || event.charCode == 45
}


////////////////////////save the path to storage by post request////////////////////
async function addNewPath() {
    if (path_names.includes(new_path_name.value)) {
        alert("pathname already exist")
    }
    else {
        new_coordinates = []
        coordinates_divs = Array.from(document.getElementsByClassName("coordinates"))
        for (div of coordinates_divs) {
            console.log(div.children)
            position = []
            for (child of Array.from(div.children)) {
                console.log(child.tagName)
                if (child.tagName.toLowerCase() == "input") {

                    position.push(parseFloat(child.value))
                }

            }
            new_coordinates.push(position)
        }

        new_path_modal.style.display = "none"
        new_path_modal.innerHTML = ''
        new_path = {
            "pathname": new_path_name.value,
            "coordinates": new_coordinates
        }


        response = await fetch('/getPath', { method: 'POST', body: JSON.stringify(new_path) })
        if (response.ok) {
            console.log(new_path)
            updatePath()

        } else {
            console.log(response)
        }
    }
}
let add_new_path_button = document.getElementById("add-new-path-button")
 
add_new_path_button.onclick = function(event){
    event.preventDefault(); 
    addNewPath()
}


//////////////////////////////////////////main function///////////////////////////
async function addGeoJson() {


    const response = await fetch(`/getPath`);
    const datas = await response.json();
    const Durations_response = await fetch(`/getDuration`)
    const Distance_response = await fetch(`/getDistance`)
    Durations = await Durations_response.json();
    Distances = await Distance_response.json();

    mypath = L.geoJson(datas, {
        filter: function (feature, layer) {
            if (path_select.value !== "all") {
                return feature.properties.pathname === path_select.value;
            }
            return true

        },
        onEachFeature: function (feature, layer) {
            for (cor of feature.geometry.coordinates) {
                markers.push(L.circleMarker([cor[1], cor[0]]).bindPopup(`
                        longitude: ${cor[0]}\n
                        latitude:${cor[1]}\n
                        altitude:${cor[0]}\n 
                        timestamp:${cor[3]}`)
                )
            }
            if (layer instanceof L.Polyline && path_select.value == "all") {
                color = "#" + Math.floor(Math.random() * 16777215).toString(16)
                layer.setStyle({
                    'color': color
                });
                legend_div.innerHTML += `<i style="background: ${color}"></i><span>${feature.properties.pathname}</span><br>`
            }
            else {
                legend_div.innerHTML += `<i style="background: blue"></i><span>${feature.properties.pathname}</span><br>`

            }

        }
    })
    addPathDataDivs()
    map.fitBounds(mypath.getBounds())

    mypath.addTo(map);
    for (marker of markers) {
        marker.addTo(map)
    }




}





//////////////////////////main////////////////////////////////////
addOptions();

addGeoJson();
