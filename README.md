# go-algorithms-basics-course

Live demo [here](https://anneauintegre.github.io/go-algorithms-basics-course)

To run locally, go in `docs/slidev-theme-shodo-main` directory. Then, install the dependencies with `npm install` and start the project with `npm run dev`.

If you don't have node, you can start a Docker container in `docs/slidev-theme-shodo-main` using the following command.

```bash
docker run --name slidev --rm -it \
    --user node \
    -v ${PWD}:/slidev \
    -p 3030:3030 \
    tangramor/slidev:latest
```