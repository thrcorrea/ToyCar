import './App.css';
import Car from './components/car';
import PlacementForm from './components/form';
import TableTop from './components/tabletop';
import React, { useEffect, useState } from 'react';

function App() {
  const [carPosition, setCarPosition] = useState({ row: null, column: null, direction: ""});

  useEffect(() => {
    const interval = setInterval(async () => {
      try {
        const response = await fetch('http://localhost:3001/toy')
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

  const handleCommand = async (command) => {
    try{
      const response = await fetch('http://localhost:3001/toy', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify({ ...command })
      })
      if (response.status == 202) {
        const data = await response.json()
        setCarPosition({row: data.y, column: data.x, direction: data.direction})
      }
    } catch (error) {
      console.error('erro ao buscar dados', error)
    }
  };

  const placeCar = (position) => {
    handleCommand({ command: "PLACE", row: position.y, column: position.x, direction: position.direction})
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
      <button onClick={() => handleCommand({ command: 'left' })}>Left</button>
      <button onClick={() => handleCommand({ command: 'right' })}>Right</button>
      <button onClick={() => handleCommand({ command: 'move' } )}>Move</button>
    </div>
      </header>
    </div>
  );
}

export default App;
