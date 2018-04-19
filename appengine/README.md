# runefinder

A simple Web app for finding Unicode characters by name. Runs on Google AppEngine.

## Local testing

Using `pipenv` to run the ancient Python 2.7 required by Google:

```
$ pipenv --two run dev_appserver.py app.yaml --enable_watching_go_path False
```

## Deploy to AppEngine and visit it there

```
$ gcloud app deploy
$ gcloud app browse
```
