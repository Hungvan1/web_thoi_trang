import React, { useContext } from 'react';
import { CartContext } from '../context/CartContext';
import CartItem from '../components/CartItem';
import { Link, useNavigate } from 'react-router-dom';

const Cart = () => {
    const { cart, updateQty, removeFromCart, total } = useContext(CartContext);
    const navigate = useNavigate();

    return (
        <div className="container cart-page">
            <h2>Giỏ hàng</h2>
            {cart.length === 0 ? (
                <div>
                    Giỏ hàng trống. <Link to="/products">Tiếp tục mua sắm</Link>
                </div>
            ) : (
                <div className="cart-grid">
                    <div className="cart-list">
                        {cart.map(it => (
                            <CartItem key={it.id} item={it} onQtyChange={updateQty} onRemove={removeFromCart} />
                        ))}
                    </div>

                    <div className="cart-summary">
                        <h3>Thanh toán</h3>
                        <p>Tổng: ${total.toFixed(2)}</p>
                        <button onClick={() => navigate('/checkout')}>Thanh toán</button>
                    </div>
                </div>
            )}
        </div>
    );
};

export default Cart;
