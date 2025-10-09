import React from 'react';

const CartItem = ({ item, onQtyChange, onRemove }) => {
    return (
        <div className="cart-item">
            <img src={item.image} alt={item.title} />
            <div className="ci-info">
                <h4>{item.title}</h4>
                <p>Price: ${item.price.toFixed(2)}</p>
                <div>
                    Qty: <input type="number" min="1" value={item.quantity} onChange={(e) => onQtyChange(item.id, Number(e.target.value))} />
                    <button onClick={() => onRemove(item.id)}>Remove</button>
                </div>
            </div>
        </div>
    );
};

export default CartItem;
