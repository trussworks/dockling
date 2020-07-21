# Dockling Layers Exercise

We are going to explore packaging a simple python application.  It
will run [docklingcake/example.py](./docklingcake/example.py) which
prints an HTTP status code after fetching a URL.

## Build and run example exercise

Take a look at `./Dockerfile.exercise` for a commented description of
what the dockerfile does.

Build the image.

    docker build -t dockling:python -f Dockerfile.exercise .

Turns out that [pipenv has a bug with
--system](https://github.com/pypa/pipenv/issues/4220) so it prints out
warnings, but it all works.

What does that do?
1. `docker build` says you want to build a image.
1. `-t dockling:python` says you want to name your image `dockling` and
   you want the name of the tag to be `python`.  If you didn't provide
   anything after the colon (e.g. `docker build -t dockling .`) that
   would create the `dockling` image with the tag `latest`.
1. `-f Dockerfile.exercise` says to use `Dockerfile.exercise` as the
   build file instead of the default `Dockerfile`.
1. `.` Means use the current directory as the context when building
   the image.  This means in the `Dockerfile` when you copy a file
   from `.` it will use that directory.

Run

    docker run --rm dockling:python

What does that do?
1. `docker run` says you want to run a new container
2. `--rm` says after this container exits, you want to remove it.
3. `dockling:python`


Now, build another tag.

Run

    docker build -t dockling:python-try-2 -f Dockerfile.exercise .

Notice that it's really fast the 2nd time and you see `Using cache`
after each Step.  Because nothing in the image has changed, it can
use the cache.

Also notice that each instruction in the `Dockerfile.exercise` is a step.  Each
one of those is called a "layer".

## What happens if you make a change

Now pretend that you want to make a local `TODO.txt` file for your own
personal use.

    echo '1. My first todo' > TODO.txt
    
Now, let's rebuild our docker image.

    docker build -t dockling:python -f Dockerfile.exercise .

Notice that the cache isn't used any more.  This is because the
`Dockerfile.exercise` does `COPY . .` so if *anything* in the current directory
changes, the docker image will have to be rebuilt.

The build process will be must faster if we can ensure we only copy
the files that are needed by the package.

## Trying the optimized build

Take a look at `./Dockerfile.optimized` for a commented description of
what the dockerfile does.

    docker build -t dockling:python -f Dockerfile.optimized .

You've replaced the `dockling:python` tag by building from the
optimized dockerfile.  Now pretend you add another todo item.

    echo '2. My second todo' >> TODO.txt
    
Now, let's rebuild our docker image.

    docker build -t dockling:python -f Dockerfile.optimized .

Notice that the cache is still used!

## Changing the code

Let's change the code.  Edit `docklingcake/example.py` and change the
URL to `http://example.com`.

Now, let's rebuild our docker image.

    docker build -t dockling:python -f Dockerfile.optimized .

If you scroll up, you can see that the cache is used for installing
pipenv and the dependencies.  The only thing that has to be run again
is `python setup.py install` so the build is *much* faster than
redoing everything.

## Bonus homework

Try `docker images` to see the images you have on your machine.
Change some files and rebuild images.  How does it change?  Try
`docker rmi some_image_tag`.
