import './App.css';
import Car from './components/car';
import PlacementForm from './components/form';
import TableTop from './components/tabletop';
import React, { useEffect, useState } from 'react';

function App() {
  // @ts-ignore
  const API_URL = process.env.REACT_APP_API_URL
  const [carPosition, setCarPosition] = useState({ row: null, column: null, direction: ""});

  useEffect(() => {
    const interval = setInterval(async () => {
      try {
        const response = await fetch(API_URL + '/toy')
        if (response.status == 200) {
          const data = await response.json()
          setCarPosition({row: data.y, column: data.x, direction: data.direction})
        }
      } catch (error) {
        console.error('erro ao buscar dados', error)
      }
    }, 5000);

    return () => clearInterval(interval)
  }, []);

  const handleCommand = async (command, data) => {
    try{
      const url = API_URL + '/toy/' + command
      const requestData = {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
      }
      if (data) {
        requestData.body = JSON.stringify(data)
      }

      const response = await fetch(url, requestData)
      if (response.status == 202) {
        const data = await response.json()
        setCarPosition({row: data.y, column: data.x, direction: data.direction})
      }
    } catch (error) {
      console.error('erro ao buscar dados', error)
    }
  };

  const placeCar = (position) => {
    handleCommand("place", position)
  }


  return (
    <div className="App">
      <header className="App-header">
      <TableTop>
        {(row, column) => {
          if (carPosition.row != null) {
            if (row === carPosition.row && column === carPosition.column) {
              return <Car {...carPosition} />;
            }
          }
          return null;
        }}
      </TableTop>
      <div>
        <PlacementForm placeCar={placeCar} carPosition={carPosition}></PlacementForm>
      <button onClick={() => handleCommand('left')}>Left</button>
      <button onClick={() => handleCommand('right')}>Right</button>
      <button onClick={() => handleCommand('move')}>Move</button>
    </div>
      </header>
    </div>
  );
}

export default App;
