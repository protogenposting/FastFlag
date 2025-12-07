# Fast Flag package manager cheet sheet

Fast Flag is a layer ontop of portage that helps with easily installing and configuring packages while not sacrificing the power of portage!

`www-client/firefox` is the package that will be used as the main example here. Replace it with whatever actual package you want!

# Basics of package management

## Basic package installation

Packages can be installed simply like this:

`fastflag install www-client/firefox`

This will simply install the package with all of the default flags. Nothing else.

This is equivelant to `emerge -av www-client/firefox`

## Basic package removal

Sometimes you messed up and want to remove a package. This is how you do that!

To remove a package that nothing else depends on, use this command:

`fastflag remove www-client/firefox`

This is equivelant to `emerge --ask --verbose --depclean www-client/firefox`

# Basics of USE flags

USE flags are what let you disable and enable certain features of the OS and individual packages. This is one of the most important components of Gentoo and Fast Flag!!!

## Find all available flags for a package

Before messing with any flags, you're going to want to know what's available.

To find every flag for a package, use this command:

`fastflag flag find www-client/firefox`

This is equivelant to `equery uses www-client/firefox`

## Changing flags

Now it's time to enable some flags! The following command enables X11 support for firefox.

`fastflag flags change www-client/firefox +X`

The + before the flag name indicates we are enabling it. To disable the X11 flag, use the command like this:

`fastflag flags change www-client/firefox -X`

You can also do multiple changes in one command like this:

`fastflag flags change www-client/firefox +X -gnome-shell -jack`

this will enable the X11 flag alongside disabling the gnome-shell and jack flags

## Viewing enabled/disabled flags

You will probably want to see what flags you enabled/disabled before updating the package. To do this, use:

`fastflag flags status www-client/firefox`

This will list all flags enabled and disabled for the specified package

## Update package with new flags

Changing flags doesn't immediatly update the package, you need to re-emerge the package first.

`fastflag flag update www-client/firefox`

This is equivelant to `emerge --update --deep --newuse www-client/firefox`

You may also want to update multiple packages. For that, use this command:

`fastflag flag update @world`
