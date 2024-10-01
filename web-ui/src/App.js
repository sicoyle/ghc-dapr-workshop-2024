import React, { useState } from 'react';

const App = () => {
  const [gameID, setGameID] = useState('');
  const [data, setData] = useState(null);

  const fetchScoreboard = () => {
    const url = process.env.REACT_APP_GAME_SERVICE_URL + "/score/"
    console.log("Calling: " + url + gameID)
    fetch(url + gameID)
      .then(response => response.json())
      .then(data => {
        setData(data);
      })
      .catch(error => console.error(error));
  };

  return (
    <>
      <div className="navbar-container">
        <nav className="navbar">
          <div className="navbar-item">
            <a href="https://dapr.io/" target="_blank" rel="noopener noreferrer">Dapr</a>
            <a href={process.env.REACT_APP_ZIPKIN_URL}>Zipkin</a>
          </div>
        </nav>
      </div>

      <div className="welcome">
        <p>Welcome to our Dapr Volleyball Demo!</p>
      </div>


      <div className="game-id-container">
        <label htmlFor="gameID">Enter Game ID:</label>
        <input type="text" id="gameID" value={gameID} onChange={e => setGameID(e.target.value)} />
        <button onClick={fetchScoreboard}>Get Game Score</button>
      </div>
      {data && (
        <div>
          <h1>Final Game Score</h1>
          <h2>{JSON.parse(data).firstTeamName} vs {JSON.parse(data).secondTeamName}</h2>
          <p>{JSON.parse(data).firstTeamName} score: {JSON.parse(data).firstTeamScore}</p>
          <p>{JSON.parse(data).secondTeamName} score: {JSON.parse(data).secondTeamScore}</p>
        </div>
      )}
    </>
  );
};

export default App;