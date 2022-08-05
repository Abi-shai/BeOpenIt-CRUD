import './film-cart.scss';


import { useState } from 'react';


const FilmCart = () => {
    // TODO: implement imageURl
    const {imageUrl, setImageUrl} = useState("");
    return (
        <div className="film">
            <img
                className='film_image'
                src="https://lestoilesfilantes.org/wp-content/uploads/Affiche-Spiderman-NoWayHome.jpeg"
            />
            <div className="film_infos">
                <span><span className='film_infos_name'>Nom:</span> Insert</span><br/>
                <span><span className='film_infos_name'>Catégorie:</span> Insert</span><br/>
            </div>
        </div>
    );
};

export default FilmCart;