import FilmCart from "../film-cart/film-cart.component";

import './films-container.style.scss';

// TODO: Implement Http route for retreving film data
const FilmsContainer = () => {
    let fakeFilms = [0, 1, 2, 3, 5];
    return (
        <div className="films_container">
            {fakeFilms.map((film, i) => {
                return <FilmCart key={i}/>
            })}
        </div>
    );
};

export default FilmsContainer;