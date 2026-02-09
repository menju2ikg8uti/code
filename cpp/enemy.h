#pragma once
#include <SFML/Graphics.hpp>
class Player;

class Character {
public:
    virtual ~Character() = default;
    virtual void update(float dt) = 0;
    virtual void draw(sf::RenderWindow& win) = 0;
    virtual sf::FloatRect bounds() const = 0;
    virtual bool alive() const { return alive_; }
protected:
    bool alive_ = true;
};

class Enemy : public Character {
public:
    virtual void onStomped(Player& p);
    virtual void onTouch(Player& p);
    virtual sf::Vector2f center() const = 0;

    sf::Font font;
    sf::Text label; // green floating text
};
