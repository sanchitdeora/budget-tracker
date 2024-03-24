import React from 'react'

function FilterButton({button, filter}) {
    return (
        <div className="filter-buttons-container">
            {
                button.map((cat, i)=>{
                    return <button type="button" onClick={()=> filter(cat)} className="filter-buttons">{cat}</button>
                })
            }
        </div>
    )
}

export default FilterButton;