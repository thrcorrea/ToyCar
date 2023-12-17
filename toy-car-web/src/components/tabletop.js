import React from 'react'
import './tabletop.css'

function TableTop({ children }) {
  const gridSize = 5;
  const rows = Array.from({ length: gridSize }, (_, rowIndex) => (
    <tr key={rowIndex} className="row">
      {Array.from({ length: gridSize }, (_, cellIndex) => (
        <td key={cellIndex} className="cell">{children(rowIndex, cellIndex)}</td>
      ))}
    </tr>
  ));

  return( 
    <div className='table-container'>
      <table className="tabletop" border={1} >
        <tbody>
          {rows}
        </tbody>
      </table>
    </div>
  );
}

export default TableTop
