import './App.scss';

import BarreDeRecherche from './components/search-bar/search-bar.component';
import FilmsContainer from './components/fims-container/films-container.component';


function App() {
  return (
    <div className='app'>
      <p className='getGameLine'>Retrouve ton film préféré <span className='exclamation-point'>!</span></p>
      <BarreDeRecherche />
      <FilmsContainer />
    </div>
  );
}

export default App;
