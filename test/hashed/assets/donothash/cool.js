earches existing DOM for elements of our component type and upgrades them
* if they have not already been upgraded.
*
* @param {string=} optJsClass the programatic name of the element class we
* need to create a new instance of.
* @param {string=} optCssClass the name of the CSS class elements of this
* type will have.
*/
function upgradeDomInternal(optJsClass, optCssClass) {
    if (typeof optJsClass === 'undefined' &&
        typeof optCssClass === 'undefined') {
        for (var i = 0; i < registeredComponents_.length; i++) {
            upgradeDomInternal(registeredComponents_[i].className,
                registeredComponents_[i].cssClass);
        }
    } else {
        var jsClass = /** @type {string} */ (optJsClass);
        if (typeof optCssClass === 'undefined') {
            var registeredClass = findRegisteredClass_(jsClass);
            if (registeredClass) {
                optCssClass = registeredClass.cssClass;
            }
        }

        var elements = document.querySelectorAll('.' + optCssClass);
        for (var n = 0; n < elements.length; n++) {
            upgradeElementInternal(elements[n], jsClass);
        }
    }
}

/**
 * Upgrades a specific element rather than all in the DOM.
*