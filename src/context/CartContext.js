import React, { createContext, useState, useEffect } from 'react';

export const CartContext = createContext();

export const CartProvider = ({ children }) => {
    const [cart, setCart] = useState(() => {
        try {
            return JSON.parse(localStorage.getItem('cart')) || [];
        } catch { return []; }
    });

    useEffect(() => {
        localStorage.setItem('cart', JSON.stringify(cart));
    }, [cart]);

    const addToCart = (product, qty = 1, selected = {}) => {
        setCart(prev => {
            const idx = prev.findIndex(i => i.id === product.id && JSON.stringify(i.selected) === JSON.stringify(selected));
            if (idx >= 0) {
                const copy = [...prev];
                copy[idx].quantity += qty;
                return copy;
            }
            return [...prev, { ...product, quantity: qty, selected }];
        });
    };

    const updateQty = (id, qty) => {
        setCart(prev => prev.map(i => i.id === id ? { ...i, quantity: qty } : i));
    };

    const removeFromCart = (id) => {
        setCart(prev => prev.filter(i => i.id !== id));
    };

    const clearCart = () => setCart([]);

    const total = cart.reduce((s, it) => s + (it.price * it.quantity), 0);

    return (
        <CartContext.Provider value={{ cart, addToCart, updateQty, removeFromCart, clearCart, total }}>
            {children}
        </CartContext.Provider>
    );
};
