import './search-bar.scss';

const searchBar = () => {
    return(
        <div className='searchBar_wrapper'>
            <form className='searchBar'>
                <input className="searchBar_input" type="text">
                </input>
                <button className='searchBar_button'>Rechercher</button>
            </form>
        </div>
    )
};

export default searchBar;