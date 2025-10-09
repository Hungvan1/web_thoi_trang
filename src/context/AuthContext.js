import React, { createContext, useState, useEffect } from 'react';

export const AuthContext = createContext();

export const AuthProvider = ({ children }) => {
    const [user, setUser] = useState(null);

    useEffect(() => {
        const saved = localStorage.getItem("user");
        if (saved) setUser(JSON.parse(saved));
    }, []);

    const login = (username, password) => {
        // tai khoan Admin
        if (username === "admin") {
            if (password === "admin123") {
                const adminUser = { username: "admin", role: "admin" };
                setUser(adminUser);
                localStorage.setItem("user", JSON.stringify(adminUser));
                return { success: true, role: "admin" };
            } else {
                return { success: false, message: "WRONG_PASSWORD" };
            }
        }

        // user thường (ví dụ)
        if (username === "testuser") {
            if (password === "test123") {
                const fakeUser = { username, role: "user" };
                setUser(fakeUser);
                localStorage.setItem("user", JSON.stringify(fakeUser));
                return { success: true, role: "user" };
            } else {
                return { success: false, message: "WRONG_PASSWORD" };
            }
        }

        return { success: false, message: "USERNAME_NOT_FOUND" };
    };

    const register = (username, password) => {
        const newUser = { username, role: "user" };
        setUser(newUser);
        localStorage.setItem("user", JSON.stringify(newUser));
    };

    const logout = () => {
        setUser(null);
        localStorage.removeItem("user");
    };

    return (
        <AuthContext.Provider value={{ user, login, register, logout }}>
            {children}
        </AuthContext.Provider>
    );
};
